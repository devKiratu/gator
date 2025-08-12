package config

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/devKiratu/gator/internal/database"
	"github.com/google/uuid"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type state struct {
	db     *database.Queries
	config *Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	Data map[string]func(*state, command) error
}

func getCommands() commands {
	return commands{
		Data: map[string]func(*state, command) error{
			"login":    handlerLogin,
			"register": handlerRegister,
		},
	}
}

func (c *commands) run(s *state, cmd command) error {
	found, ok := c.Data[cmd.name]
	if !ok {
		return fmt.Errorf("unkown command: %s", cmd.name)
	}
	err := found(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Data[name] = f
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	c := Config{}
	file, err := getConfigFilePath()
	if err != nil {
		return c, err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := home + "/" + configFileName
	return filePath, nil
}

/* Command Handlers */
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("a username is required")
	}
	if len(cmd.args) > 3 {
		return fmt.Errorf("login command expects a single argument, the username")
	}
	// check if user exists in db
	user, err := s.db.GetUser(context.Background(), cmd.args[2])
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	if user.Name == "" {
		return fmt.Errorf("account %s doesn't exist", cmd.args[2])
	}
	err = s.config.SetUser(cmd.args[2])
	if err != nil {
		return fmt.Errorf("error setting current user: %w", err)
	}
	fmt.Println("user set successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("a username is required")
	}
	if len(cmd.args) > 3 {
		return fmt.Errorf("login command expects a single argument, the username")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[2]})
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}
	err = s.config.SetUser(cmd.args[2])
	if err != nil {
		return fmt.Errorf("error setting current user: %w", err)
	}
	fmt.Printf("User created: %+v\n", user)
	return nil
}

func StartGator() {
	data, err := Read()
	if err != nil {
		fmt.Println("Error reading file: %w", err)
	}
	//db connection
	db, err := sql.Open("postgres", data.DbUrl)
	if err != nil {
		fmt.Println("Error establishing db connection", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	// fmt.Printf("BEFORE: %+v", data)
	state := state{
		db:     dbQueries,
		config: &data,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	currentCommand := strings.ToLower(args[1])
	commands := getCommands()
	err = commands.run(&state, command{name: currentCommand, args: args})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Printf("AFTER: %+v", state.config)
}
