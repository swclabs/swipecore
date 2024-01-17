// Authors: Duc Hung Ho @kieranhoo
// Description: Plugin for tasks package, run worker tasks in memory

package resolver

import (
	"log"
	"mime/multipart"
	"sync"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/repo"
	"github.com/swclabs/swipe-api/pkg/cloud"
)

var ImagePool *Pool

func StartImageHandler(concurrent int) {
	ImagePool = NewPool()
	img := Image{}
	NewTask(ImagePool, img.ImageHandler, concurrent).Run()
}

type Image struct {
	Email string
	File  multipart.File
}

func (image *Image) ImageHandler(data <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _data := range data {
		if imageInfo, ok := _data.(Image); ok {
			resp, err := cloud.UpdateImages(cloud.Connection(), imageInfo.File)
			if err != nil {
				log.Fatal(err)
			}
			if err := repo.NewUsers().SaveInfo(&domain.User{
				Email: imageInfo.Email,
				Image: resp.SecureURL,
			}); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("can't update images: unknown image type")
		}
	}
}