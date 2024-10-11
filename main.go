package main

import (
    "blog/photos-blog/cloudinary"
)

func main() {
    cloudinary.Init()
    cloudinary.GetAllImagesInFolder("photos_blog")
}
