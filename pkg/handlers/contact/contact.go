package contactHandler

import (
	"net/http"

	contactRepo "github.com/IlyesDEO/goCrud/pkg/repositories/contact"
	"github.com/gin-gonic/gin"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func GetAll(c *gin.Context) {
	var (
		contacts []*structs.Contact
		err      error
	)

	if contacts, err = contactRepo.GetAll(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot retrieve data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func GetById(c *gin.Context) {
	var (
		contact *structs.Contact
		err     error
	)

	if contact, err = contactRepo.GetById(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot retrieve data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func Create(c *gin.Context) {
	var (
		contact *structs.Contact
		err     error
	)

	if err = c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot bind data: " + err.Error()})
		return
	}

	if err = contactRepo.Create(contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot create data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func Update(c *gin.Context) {
	var (
		contact *structs.Contact
		err     error
	)

	if err = c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot bind data: " + err.Error()})
		return
	}

	if err = contactRepo.Update(contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot update data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func Delete(c *gin.Context) {
	var (
		err error
	)

	if err = contactRepo.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot delete data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": c.Param("id")})
}
