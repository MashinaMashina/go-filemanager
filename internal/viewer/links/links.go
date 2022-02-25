package links

import (
	"net/url"
)

func DownloadFile(file string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", file)
	query.Add("action", "download")
	u.RawQuery = query.Encode()

	return u
}

func CreateFile(dir string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", dir)
	query.Add("action", "create_file")
	u.RawQuery = query.Encode()

	return u
}

func CreateDir(dir string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", dir)
	query.Add("action", "create_dir")
	u.RawQuery = query.Encode()

	return u
}

func UploadFile(dir string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", dir)
	query.Add("action", "upload")
	u.RawQuery = query.Encode()

	return u
}

func View(path string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", path)
	u.RawQuery = query.Encode()

	return u
}

func Rename(path string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", path)
	query.Add("action", "rename")
	u.RawQuery = query.Encode()

	return u
}

func Remove(path, csrf string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", path)
	query.Add("action", "remove")
	query.Add("_csrf", csrf)
	u.RawQuery = query.Encode()

	return u
}
