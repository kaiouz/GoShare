package res

import (
	"fmt"
	"github.com/satori/go.uuid"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

// 资源类型
type ResourceType string

const (
	Image ResourceType = "image" // 图片
	Audio ResourceType = "audio" // 音频
	Video ResourceType = "video" // 视频
)

// 资源
type Resource struct {
	Id        string       `json:"id"`
	Name      string       `json:"name"`
	ResType   ResourceType `json:"resType"`
	Thumbnail string       `json:"thumbnail"`
	Path      string       `json:"path"`
}

// 读取的资源列表
var resources []Resource

// 扫描目录中的文件生成资源
func CreateResources(dirs []string) []Resource {
	resources = make([]Resource, 0, 30)
	for _, dir := range dirs {
		if err := filepath.Walk(dir, walk); err != nil {
			fmt.Println(err)
		}
	}
	return resources
}

func walk(path string, info os.FileInfo, err error) error {
	fmt.Printf("处理: %s\n", path)
	// 错误跳过这个文件
	if err != nil {
		fmt.Printf("跳过文件[%s], 因为发生错误: %v\n", path, err)
		return filepath.SkipDir
	}

	// 目录无需处理
	if info.IsDir() {
		return nil
	}

	// 加入创建好的资源
	if res, err := createResource(path, info); err == nil {
		resources = append(resources, res)
	}

	return err
}

// 从文件生成资源
func createResource(path string, info os.FileInfo) (Resource, error) {
	res := Resource{
		Id:   uuid.Must(uuid.NewV4()).String(),
		Path: path,
	}

	ext := filepath.Ext(path)
	mimeType := mime.TypeByExtension(ext)

	// 文件名
	res.Name = strings.TrimSuffix(info.Name(), ext)

	switch {
	case strings.HasPrefix(mimeType, string(Image)):
		res.ResType = Image
		return createImageRes(res)
	case strings.HasPrefix(mimeType, string(Audio)):
		res.ResType = Audio
		return createAudioRes(res)
	case strings.HasPrefix(mimeType, string(Video)):
		res.ResType = Video
		return createVideoRes(res)
	}

	// 不支持的文件跳过
	return res, filepath.SkipDir
}

// 创建视频资源
func createVideoRes(resource Resource) (Resource, error) {
	return resource, nil
}

// 创建音频资源
func createAudioRes(resource Resource) (Resource, error) {
	return resource, nil
}

// 创建图片资源
func createImageRes(resource Resource) (Resource, error) {
	return resource, nil
}
