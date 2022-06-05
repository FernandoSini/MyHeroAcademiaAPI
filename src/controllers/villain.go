package controllers

import (
	"MyHeroAcademiaApi/src/database"
	"MyHeroAcademiaApi/src/models"
	"MyHeroAcademiaApi/src/repository"
	"MyHeroAcademiaApi/src/responses"
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllVillains(w http.ResponseWriter, _ *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewVillainRepository(db)
	villains, err := repo.FindVillains()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if len(villains) <= 0 && err == nil {
		responses.Erro(w, http.StatusNotFound, errors.New(" Not found villains! "))
		return
	}

	responses.JSON(w, http.StatusOK, villains)

}

func FindVillainById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	villainId := params["villainId"]
	if len(villainId) <= 0 || villainId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New(" Not found this villain! "))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())
	repo := repository.NewVillainRepository(db)
	villain, err := repo.FindVillainByID(villainId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if villain.Id.Hex() != villainId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}

	responses.JSON(w, http.StatusOK, villain)
}

/* func CreateVillain(w http.ResponseWriter, r *http.Request) {

	///getting the body of requisiton and converting to json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var villain models.Villain
	villain.Id = primitive.NewObjectID()
	villain.Thumbnail.Id = primitive.NewObjectID()
	villain.Thumbnail.IdVillainRef = villain.Id
	//getting the json sent by the reqBody and converting to villain model
	//marshal --> converts byte/data to json
	//unmarshal -> converts json to model data
	if err = json.Unmarshal(reqBody, &villain); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = villain.Preparar(); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewVillainRepository(db)
	err = repo.CreateVillain(villain)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, "Villain created Successfully")

} */

func CreateVillain(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, errors.New("Can't upload image: "+err.Error()))
		return
	}

	/* b := make([]byte, r.MultipartForm.File["file"][0].Size)
	binary.LittleEndian.PutUint64(b, uint64(r.MultipartForm.File["file"][0].Size))

	dataBytes := bytes.NewBuffer(b) */

	var byteData []byte
	file, handler, err := r.FormFile("file")
	if file == nil || handler.Size == 0 {
		var villain models.Villain
		villain.Id = primitive.NewObjectID()
		villain.VillainName = r.PostFormValue("villainName")
		villain.LastName = r.PostFormValue("lastName")
		villain.TrueName = r.PostFormValue("trueName")
		villain.Description = r.PostFormValue("description")
		villain.Age, err = strconv.ParseInt(r.PostFormValue("age"), 10, 64)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, errors.New("error while trying to convert string to int on age"+err.Error()))
			return
		}
		villain.VillainRank, err = strconv.ParseInt(r.PostFormValue("villainRank"), 10, 64)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, errors.New("error while trying to convert string to int on villainRank"+err.Error()))
			return
		}

		//getting the json sent by the reqBody and converting to villain model
		//marshal --> converts byte/data to json
		//unmarshal -> converts json to model data
		/* if err = json.Unmarshal(reqBody, &villain); err != nil {
			responses.Erro(w, http.StatusBadRequest, err)
			return
		} */

		if err = villain.Preparar(); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}

		db, err := database.Connect()
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Disconnect(context.Background())

		repo := repository.NewVillainRepository(db)
		err = repo.CreateVillain(villain)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, "Villain created Successfully")

	} else {
		if err != nil {
			responses.Erro(w, http.StatusBadRequest, errors.New("Problem to upload image file: "+err.Error()))
			return
		}
		defer file.Close()

		if handler.Header.Get("Content-Type") == "image/png" {

			tempFile, err := ioutil.TempFile("tmp", "file-*.png")
			if err != nil {
				responses.Erro(w, http.StatusBadRequest, errors.New("error on temp file png: "+err.Error()))
				return
			}
			defer tempFile.Close()

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				responses.Erro(w, http.StatusBadRequest, errors.New("error on filebytes png: "+err.Error()))
				return
			}
			byteData = fileBytes
			//tempFile.Write(fileBytes)

		}
		if handler.Header.Get("Content-Type") == "image/jpeg" {
			tempFile, err := ioutil.TempFile("tmp", "file-*.jpg")
			if err != nil {
				responses.Erro(w, http.StatusBadRequest, errors.New("error on temp file jpeg: "+err.Error()))
				return
			}
			defer tempFile.Close()

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				responses.Erro(w, http.StatusBadRequest, errors.New("error on filebytes jpeg: "+err.Error()))
				return
			}
			byteData = fileBytes
			//tempFile.Write(fileBytes)

		}

		var villain models.Villain
		villain.Id = primitive.NewObjectID()
		villain.VillainName = r.PostFormValue("villainName")
		if byteData != nil {
			villain.Thumbnail.Id = primitive.NewObjectID()
			villain.Thumbnail.IdVillainRef = villain.Id
			villain.Thumbnail.Content, err = UploadOnAzure(byteData, strings.ToLower(villain.VillainName))
			if err != nil {
				responses.Erro(w, http.StatusInternalServerError, err)
				return
			}
		}

		villain.LastName = r.PostFormValue("lastName")
		villain.TrueName = r.PostFormValue("trueName")
		villain.Description = r.PostFormValue("description")
		villain.Age, err = strconv.ParseInt(r.PostFormValue("age"), 10, 64)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, errors.New("error while trying to convert string to int on age"+err.Error()))
			return
		}
		villain.VillainRank, err = strconv.ParseInt(r.PostFormValue("villainRank"), 10, 64)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, errors.New("error while trying to convert string to int on villainRank"+err.Error()))
			return
		}

		//getting the json sent by the reqBody and converting to villain model
		//marshal --> converts byte/data to json
		//unmarshal -> converts json to model data
		/* if err = json.Unmarshal(reqBody, &villain); err != nil {
			responses.Erro(w, http.StatusBadRequest, err)
			return
		} */

		if err = villain.Preparar(); err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}

		db, err := database.Connect()
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Disconnect(context.Background())

		repo := repository.NewVillainRepository(db)
		err = repo.CreateVillain(villain)
		if err != nil {
			responses.Erro(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, "Villain created Successfully")
		//responses.JSON(w, 200, dataBytes.Bytes())
	}
}

func UploadOnAzure(bytesToUpload []byte, villainName string) (string, error) {

	url, ok := os.LookupEnv("AZURE_URL")
	if !ok {
		panic(errors.New("AZURE_URL could not be found"))
	}

	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic(errors.New("AZURE_STORAGE_ACCOUNT_NAME could not be found"))
	}
	accountKey, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_KEY")
	if !ok {
		panic(errors.New("AZURE_STORAGE_ACCOUNT_KEY could not be found"))
	}

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, errors.New("Invalid credentials with error: "+err.Error()))
		return "", errors.New("Invalid credentials with error: " + err.Error())
	}

	serviceClient, err := azblob.NewServiceClientWithSharedKey(url, credential, nil)
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, errors.New("Invalid credentials with error: "+err.Error()))
		return "", errors.New("Invalid credentials with error: " + err.Error())
	}

	containerName := string(os.Getenv("CONTAINER_NAME"))
	fmt.Printf("Creating container %s \n", containerName)

	containerClient, err := serviceClient.NewContainerClient(containerName + villainName)
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, errors.New("Error on container "+err.Error()))
		return "", errors.New("Error on container " + err.Error())
	}

	_, err = containerClient.Create(context.Background(), &azblob.ContainerCreateOptions{Access: azblob.PublicAccessTypeBlob.ToPtr()})
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, err)
		return "", errors.New("error 1234" + err.Error())
	}

	data := bytesToUpload

	blobName := "blob-" + villainName

	blobClient, err := containerClient.NewBlockBlobClient(blobName)

	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, err)
		return "", err
	}
	_, err = blobClient.UploadBuffer(context.Background(), data, azblob.UploadOption{})

	if err != nil {
		//responses.JSON(w, http.StatusInternalServerError, errors.New("Failed to upload "+err.Error()))
		return "", errors.New("Failed to upload " + err.Error())
	}

	pager := containerClient.ListBlobsFlat(nil)

	for pager.NextPage(context.Background()) {
		resp := pager.PageResponse()

		for _, v := range resp.ListBlobsFlatSegmentResponse.Segment.BlobItems {
			fmt.Println(*v.Name)
		}
	}

	if err = pager.Err(); err != nil {
		//responses.Erro(w, http.StatusInternalServerError, errors.New("failed to list blob"+err.Error()))
		return "", errors.New("failed to list blob" + err.Error())
	}

	// Download the blob
	get, err := blobClient.Download(context.Background(), nil)
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, errors.New("error to download blob"+err.Error()))
		return "", errors.New("error to download blob" + err.Error())
	}
	downloadedData := &bytes.Buffer{}
	reader := get.Body(&azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(reader)
	if err != nil {
		//responses.Erro(w, http.StatusInternalServerError, err)
		return "", err
	}

	err = reader.Close()
	if err != nil {
		//log.Fatal(err)
		return "", err
	}
	return blobClient.URL(), nil
}

/* func CreateVillain(url string, paramsText map[string]interface{},imageVillain models.VillainFile){
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err)
	}r.MultipartForm.File["file"][0].Header.Get("Content-Type")!="image/png" essa linha pega o mimetype do arquivo
	responses.JSON(w, 200,r.MultipartForm.File["file"][0].Header.Get("Content-Type"))

} */
func UpdateVillain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	villainId := params["villainId"]
	if len(villainId) <= 0 || villainId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewVillainRepository(db)
	villainInDB, err := repo.FindVillainByID(villainId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if villainInDB.Id.Hex() != villainId {
		responses.Erro(w, http.StatusForbidden, errors.New("forbidden action"))
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var villain models.Villain
	if err = json.Unmarshal(reqBody, &villain); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = villain.Preparar(); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	/*fazer a func de preparar*/
	/*
		if erro = villain.prepare(); erro != nil {
				responses.Erro(w, http.StatusBadRequest, erro)
				return
			}
	*/

	if err = repo.UpdateVillain(villainId, villain); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

}
func DeleteVillain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	villainId := params["villainId"]
	if len(villainId) <= 0 || villainId == "" {
		responses.Erro(w, http.StatusNotFound, errors.New("not found"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewVillainRepository(db)
	if err = repo.DeleteVillain(villainId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, " Villain deleted successfully! ")

}

// func AddVillainImage(w http.ResponseWriter, r *http.Request) {
// 	const MaxUploadSize = 10 << 20
// 	params := mux.Vars(r)
// 	villainId, err := primitive.ObjectIDFromHex(params["villainId"])
// 	if err != nil {
// 		responses.Erro(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

// 	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
// 		responses.Erro(w, http.StatusNotFound, errors.New(" The uploaded file is too big. Please choose an file that's less than 1MB in size"))
// 		return
// 	}

// 	if len(villainId) <= 0 || villainId.Hex() == "" {
// 		responses.Erro(w, http.StatusInternalServerError, errors.New(" Sorry, we can't upload image to null villain"))
// 		return
// 	}

// 	file, fileHandler, err := r.FormFile("file")
// 	if err != nil {
// 		responses.Erro(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	defer file.Close()

// 	/* //creating upload directory if it doesn't exists
// 	err = os.Mkdir("./uploads", os.ModePerm) */

// 	path, err := os.Create(filepath.Join("/Users/fernandosini/Documents/go.nosync/myheroapi/MyHeroAcademiaApi/uploads", filepath.Base(fileHandler.Filename)))
// 	if err != nil {
// 		responses.Erro(w, http.StatusInternalServerError, errors.New("sorry, can't upload image to directory"))
// 		return
// 	}
// 	defer path.Close()
// 	if _, err = io.Copy(path, file); err != nil {
// 		responses.Erro(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	if _, err = io.Copy(path, file); err != nil {
// 		responses.Erro(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	//creating buffer based on file size
// 	fileInfo, _ := path.Stat()
// 	var size int64 = fileInfo.Size()
// 	buffer := make([]byte, size)

// 	//read file content into buffer
// 	fileReader := bufio.NewReader(path)
// 	fileReader.Read(buffer)

// 	//convert buffer bytes to base64 string - use buffer.Bytes() for new image
// 	imgBase64Str := base64.StdEncoding.EncodeToString(buffer)
// 	fmt.Fprintf(w, imgBase64Str)

// 	//decoding image
// 	imgStringDecoded, _ := base64.StdEncoding.DecodeString(imgBase64Str)
// 	fmt.Println(imgStringDecoded)
// 	pathFromFile := "\\MyheroAcademiaApi\\uploads\\" + fileHandler.Filename

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.Erro(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Disconnect(context.Background())

// 	repo := repository.NewVillainFileRepository(db)
// 	villainFile := models.VillainFile{}
// 	villainFile.Path = pathFromFile
// 	villainFile.FileName = fileHandler.Filename
// 	villainFile.FileData = imgBase64Str
// 	villainFile.IdVillainRef = villainId
// 	repo.AddVillainFile(villainFile)

// }

func FindVillainByVillainName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	villainName := params["villainName"]

	if len(villainName) <= 0 || villainName == "" {
		responses.Erro(w, http.StatusInternalServerError, errors.New(" Inválid villainName! "))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Disconnect(context.Background())

	repo := repository.NewVillainRepository(db)
	villain, err := repo.FindVillainByVillainName(villainName)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, errors.New(" Not found villains with this name! "))
		return
	}
	responses.JSON(w, http.StatusOK, villain)

}

func addImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, errors.New("Can't upload image"+err.Error()))
		return
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(r.MultipartForm.File["file"][0].Size))

	dataBytes := bytes.NewBuffer(b)

	responses.JSON(w, 200, dataBytes)

}
