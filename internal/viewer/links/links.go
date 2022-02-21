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
	query.Add("download", "true")
	u.RawQuery = query.Encode()

	return u
}

func CreateFile(dir string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", dir)
	query.Add("create_file", "true")
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
	query.Add("rename", "true")
	u.RawQuery = query.Encode()

	return u
}

func Remove(path, csrf string) *url.URL {
	u := &url.URL{
		Path: "/",
	}

	query := u.Query()
	query.Add("path", path)
	query.Add("remove", "true")
	query.Add("_csrf", csrf)
	u.RawQuery = query.Encode()

	return u
}
