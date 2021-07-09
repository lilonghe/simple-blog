package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"simple-blog/pkg/models"
	"time"

	"gopkg.in/yaml.v3"
)

type FileMeta struct {
	Title      string
	Date       time.Time
	Lastmod    time.Time `yaml:",omitempty`
	Draft      bool      `yaml:",omitempty"`
	Categories []string  `yaml:",flow,omitempty"`
	Tags       []string  `yaml:",flow,omitempty"`
	Private    bool      `yaml:",omitempty"`
}

var filePath = "posts"

func init() {
	models.LoadOption()
}

func main() {
	info, _ := os.Stat(filePath)
	if info != nil {
		fmt.Println("Will auto remove posts dir, if posts exists, enter continue or Ctrl + C to end.")
		fmt.Scanln()
	}

	postList := getPostList()
	// postData, err := ioutil.ReadFile("./postList.json")
	// if err != nil {
	// 	panic(err)
	// }
	// err = json.Unmarshal(postData, &postList)
	// if err != nil {
	// 	panic(err)
	// }

	err := os.RemoveAll(filePath)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, post := range postList {
		categories := make([]string, 0)
		for _, cate := range post.Cates {
			categories = append(categories, cate.Name)
		}
		tags := make([]string, 0)
		for _, tag := range post.Tags {
			tags = append(tags, tag.Name)
		}
		meta := FileMeta{
			Title:      post.Title,
			Date:       post.CreateTime,
			Lastmod:    post.UpdateTime,
			Categories: categories,
			Tags:       tags,
			Draft:      post.Status == 0,
			Private:    !post.IsPublic,
		}

		metaData, err := yaml.Marshal(meta)
		if err != nil {
			panic(err)
		}
		fileContent := "---\n" + string(metaData) + "---\n" + post.MarkdownContent
		err = ioutil.WriteFile(filePath+"/"+post.Pathname+".md", []byte(fileContent), 0666)
		if err != nil {
			panic(err)
		}
	}
	s := fmt.Sprintf("Generated %d post.\n", len(postList))
	io.WriteString(os.Stdout, s)
}

func getPostList() []models.Post {
	posts, _, err := models.GetAllPostList(99999, 0)
	if err != nil {
		panic(err)
	}

	cates, err := models.GetAllCate()
	if err != nil {
		panic(err)
	}

	tags, err := models.GetAllTag()
	if err != nil {
		panic(err)
	}

	postTags, err := models.GetAllPostTag()
	if err != nil {
		panic(err)
	}

	postCates, err := models.GetAllPostCate()
	if err != nil {
		panic(err)
	}

	for k, post := range posts {
		posts[k].Cates = getPostCates(post.Id, postCates, cates)
		posts[k].Tags = getPostTags(post.Id, postTags, tags)
	}

	return posts
}

func getPostTags(postId int32, postTags []models.PostTag, tags []models.Tag) []models.Tag {
	postTagList := make([]models.Tag, 0)
	tagIds := make([]int32, 0)
	for _, v := range postTags {
		if v.PostId == postId {
			tagIds = append(tagIds, v.TagId)
		}
	}
	for _, tagId := range tagIds {
		for _, tag := range tags {
			if tag.Id == tagId {
				postTagList = append(postTagList, tag)
				break
			}
		}
	}
	return postTagList
}

func getPostCates(postId int32, postCates []models.PostCate, cates []models.Cate) []models.Cate {
	postCateList := make([]models.Cate, 0)
	cateIds := make([]int32, 0)
	for _, v := range postCates {
		if v.PostId == postId {
			cateIds = append(cateIds, v.CateId)
		}
	}
	for _, cateId := range cateIds {
		for _, cate := range cates {
			if cate.Id == cateId {
				postCateList = append(postCateList, cate)
				break
			}
		}
	}
	return postCateList
}
