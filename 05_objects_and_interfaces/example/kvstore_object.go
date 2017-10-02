package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file := "store_simple.db"
	store, err := ReadStore(file)
	if err != nil {
		fmt.Printf("error opening store: %v", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		WriteTo(store.GetAll(), os.Stdout)
		os.Exit(0)
	}

	for _, arg := range os.Args[1:] {
		if err := handleCommand(store, arg); err != nil {
			fmt.Printf("error while handling command: %v", err.Error())
			os.Exit(2)
		}
		
	}
}

type Store interface {
	Set(key, value string) error
	Get(key string) (value string, exist bool)
	GetAll() (map[string]string)
	WriteStore() error
}

type kvstore struct {
	fileName string
	store map[string]string
}

func (kv *kvstore) Get(key string) (value string, exist bool) {
	value, exists:= kv.store[key]
	return value, exists
}

func (kv *kvstore) GetAll() (map[string]string) {
	return kv.store
}

func (kv *kvstore) Set(key, value string) error {
	kv.store[key] = value
	return kv.WriteStore()
}

func handleCommand(store Store, arg string) (error) {
	if strings.Contains(arg, "=") {

		keyvalue := strings.SplitN(arg, "=", 2)
		return store.Set(keyvalue[0], keyvalue[1])
	} else {

		if v, exist := store.Get(arg); exist {
			fmt.Printf("%v=%v\n", arg, v)
			return nil
		} else {
			fmt.Printf("%v NOT FOUND\n", arg)
			return fmt.Errorf("Could not find entry for key %s", arg)
		}
	}
}

func ReadStore(file string) (Store, error) {
	store := &kvstore{
		file,
		make(map[string]string),
	}

	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		handleCommand(store, scanner.Text())
	}

	return store, scanner.Err()
}

func (kv *kvstore) WriteStore() error {
	f, err := os.Create(kv.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return WriteTo(kv.store, f)
}

func WriteTo(store map[string]string, w io.Writer) error {
	for k, v := range store {
		if _, err := fmt.Fprintf(w, "%v=%v\n", k, v); err != nil {
			return err
		}
	}
	return nil
}