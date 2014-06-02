/* 	
	Copyright (c) 2014 @irvingprog

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	see <http://www.gnu.org/licenses/
*/

package main 

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func get_page(direccion string) string {
	resp, err := http.Get(direccion)
	if err != nil {
		fmt.Print("%s", err)
		os.Exit(1)
	} 

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Print("%s", err)
		os.Exit(1)		
	}

	return string(contents)
}

func get_text_of_tags(tag string, kclass string, html string) []string {
	tag_c := "<"+ tag +" class=\""+kclass+"\">([\\s\\w]+)|([\\d]+[.*])</"+ tag +">"
	classname, _ := regexp.Compile(tag_c)
	texts := classname.FindAllString(html, -1)

	for i := range texts {
		texts[i] = strings.Replace(texts[i], "<"+ tag +" class=\""+kclass+"\">", 
								   "", -1)
		texts[i] = strings.Replace(texts[i], "</"+ tag +">", "", -1)
	}

	return texts
}

func main() {
	var song string

	fmt.Print("Enter name song, band or artist: ")
	fmt.Scanf("%s", &song)

	html := get_page("http://www.goear.com/search/"+song)
	
	songs_name := get_text_of_tags("span", "song", html)
	songs_artist := get_text_of_tags("span", "group", html)
	songs_quality := get_text_of_tags("li", "kbps radius_3", html)

	if len(songs_name) == 0 {
		fmt.Println("\n No results \n")
	} else {
		for i := range songs_name {
			fmt.Printf("Song %d \n 	name: %v\n 	Artist: %v\n 	Quality: %v kbps\n", 
					   i+1, songs_name[i], songs_artist[i], songs_quality[i])
		}
	}
}