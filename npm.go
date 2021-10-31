package main

type NpmAuthor struct {
	name string `json: name`
}

type NpmPackage struct {
	id          string            `json: _id`
	name        string            `json: name`
	distTags    map[string]string `json: dist-tags`
	license     string            `json: license`
	author      NpmAuthor         `json: author`
	homepage    string            `json: homepage`
	description string            `json: description`
}
