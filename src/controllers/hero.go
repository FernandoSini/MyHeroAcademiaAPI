package controllers

import (
	"MyHeroAcademiaApi/src/database"
	"MyHeroAcademiaApi/src/models"
	"MyHeroAcademiaApi/src/repository"
	"MyHeroAcademiaApi/src/responses"
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllHeroes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	heroes, err := repo.FindHeroes()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if len(heroes) <= 0 && err == nil {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	responses.JSON(w, http.StatusOK, heroes)

}

func FindHeroById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	heroId := params["heroId"]
	if len(heroId) <= 0 || heroId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	hero, err := repo.FindHeroByID(heroId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if hero.Id.Hex() != heroId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}

	responses.JSON(w, http.StatusOK, hero)
}

func CreateHero(w http.ResponseWriter, r *http.Request) {

	//getting the body of requisiton and converting to json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var hero models.Hero
	//getting the json sent by the reqBody and converting to hero model
	//marshal --> converts byte/data to json
	//unmarshal -> converts json to model data
	if err = json.Unmarshal(reqBody, &hero); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	err = repo.CreateHero(hero)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "Hero created Successfully")

}
func UpdateHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	heroId := params["userId"]
	if len(heroId) > 0 || heroId != "" {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
	}

	db, err := database.Connect()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewHeroRepository(db)
	heroInDB, err := repo.FindHeroByID(heroId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	if heroInDB.Id.String() != heroId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var hero models.Hero
	if err = json.Unmarshal(reqBody, &hero); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}

	/*fazer a func de preparar*/
	/*
		if erro = hero.prepare(); erro != nil {
				responses.Erro(w, http.StatusBadRequest, erro)
				return
			}
	*/

	if err = repo.UpdateHero(heroId, hero); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

}
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	heroId := params["heroId"]
	if len(heroId) <= 0 || heroId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewHeroRepository(db)
	if err = repo.DeleteHero(heroId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)

}

func AddHeroImage(w http.ResponseWriter, r *http.Request) {
	const MaxUploadSize = 10 << 20
	params := mux.Vars(r)
	heroId, err := primitive.ObjectIDFromHex(params["heroId"])
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		responses.Erro(w, http.StatusNotFound, errors.New("The uploaded file is too big. Please choose an file that's less than 1MB in size"))
		return
	}

	if len(heroId) <= 0 || heroId.Hex() == "" {
		responses.Erro(w, http.StatusInternalServerError, errors.New(" sorry, we can't upload image to null hero"))
		return
	}

	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	/* //creating upload directory if it doesn't exists
	err = os.Mkdir("./uploads", os.ModePerm) */

	path, err := os.Create(filepath.Join("/Users/fernandosini/Documents/go.nosync/myheroapi/MyHeroAcademiaApi/uploads", filepath.Base(fileHandler.Filename)))
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New("sorry, can't upload image to directory"))
		return
	}
	defer path.Close()
	if _, err = io.Copy(path, file); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if _, err = io.Copy(path, file); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	//creating buffer based on file size
	fileInfo, _ := path.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)

	//read file content into buffer
	fileReader := bufio.NewReader(path)
	fileReader.Read(buffer)

	//convert buffer bytes to base64 string - use buffer.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buffer)
	fmt.Fprintf(w, imgBase64Str)

	//decoding image
	imgStringDecoded, _ := base64.StdEncoding.DecodeString(imgBase64Str)
	fmt.Println(imgStringDecoded)
	pathFromFile := "\\MyheroAcademiaApi\\uploads\\" + fileHandler.Filename

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewHeroImageRepository(db)
	heroImage := models.HeroImage{}
	heroImage.ImagePath = pathFromFile
	heroImage.FileName = fileHandler.Filename
	heroImage.ImgData = imgBase64Str
	heroImage.IdHeroRef = heroId
	repo.AddHeroImage(heroImage)

}
