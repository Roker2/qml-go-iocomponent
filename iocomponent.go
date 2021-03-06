/*
 * Copyright (C) 2020  Dmitry Minko
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * MEGA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package iocomponent

import (
	"github.com/nanu-c/qml-go"
	"io/ioutil"
	"log"
	"os"
)

type folderType int

const (
	Config folderType = 0
	Cache folderType = 1
	AppData folderType = 2
)

type IOComponent struct {
	configFolder string
	cacheFolder string
	appDataFolder string
}

func (io *IOComponent) getFolder(ft folderType) string {
	switch ft {
	case Config:
		return io.configFolder
	case Cache:
		return io.cacheFolder
	case AppData:
		return io.appDataFolder
	}
	return ""
}

func (io *IOComponent) WriteToFile(ft folderType, fileName string, text string) error {
	folder := io.getFolder(ft)
	if !io.IsExist(ft, "") {
		err := io.Mkdir(ft, "")
		if err != nil {
			log.Println(err)
			return err
		}
	}
	file, err := os.Create(folder + fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (io *IOComponent) ReadFromFile(ft folderType, fileName string) string {
	folder := io.getFolder(ft)
	if !io.IsExist(ft, fileName) {
		log.Println("File" + fileName + "does not exist")
		return ""
	}
	b, err := ioutil.ReadFile(folder + fileName)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

func (io *IOComponent) IsExist(ft folderType, fileName string) bool {
	folder := io.getFolder(ft)
	_, err := os.Stat(folder + fileName)
	if err != nil {
		return false
	}
	return true
}

func (io *IOComponent) CreateFile(ft folderType, fileName string) {
	folder := io.getFolder(ft)
	if !io.IsExist(ft, "") {
		err := io.Mkdir(ft, "")
		if err != nil {
			log.Println(err)
		}
	}
	file, err := os.Create(folder + fileName)
	if err != nil {
		log.Println(err)
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
	}
}

func (io *IOComponent) Mkdir(ft folderType, folder string) error {
	err := os.Mkdir(folder, os.ModeDir)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (io *IOComponent) RemoveFile(ft folderType, fileName string) {
	err := os.Remove(io.getFolder(ft) + fileName)
	if err != nil {
		log.Println(err)
	}
}

func Register(packageName string) {
	qml.RegisterTypes("GoIOComponent", 0, 2, []qml.TypeSpec{{
		Init: func(v *IOComponent, obj qml.Object) {
			log.Println("Init IOComponent")
			home, err := os.UserHomeDir()
			if err != nil {
				log.Println("GoIOComponent error: " + err.Error())
			}
			v.configFolder =  home + "/.config/" + packageName + "/"
			v.cacheFolder =  home + "/.cache/" + packageName + "/"
			v.appDataFolder =  home + "/.local/share/" + packageName + "/"
		},
	}})
}

func (io *IOComponent) SetPackageName(packageName string) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("GoIOComponent error: " + err.Error())
	}
	io.configFolder =  home + "/.config/" + packageName + "/"
	io.cacheFolder =  home + "/.cache/" + packageName + "/"
	io.appDataFolder =  home + "/.local/share/" + packageName + "/"
}
