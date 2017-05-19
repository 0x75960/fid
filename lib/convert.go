package fid

import "fmt"

// MarkdownTable text from FileIdentifer
func (fid *FileIdentifer) MarkdownTable() (str string) {

	str = fmt.Sprintf(`| Item   | Value |
| :----  | :---- |
| Path   | %s    |
| SHA1   | %s    |
| SHA256 | %s    |
| MD5    | %s    |
`, fid.Path, fid.Sha1, fid.Sha256, fid.Md5)

	return
}

// MarkdownTableTitle text
func MarkdownTableTitle() (str string) {
	return "|Path|SHA1|SHA256|MD5|\n|:--|:--|:--|:--|\n"
}

// MarkdownTableRow text from FileIdentifer
func (fid *FileIdentifer) MarkdownTableRow() (str string) {
	return fmt.Sprintf("|%s|%s|%s|%s|\n", fid.Path, fid.Sha1, fid.Sha256, fid.Md5)
}
