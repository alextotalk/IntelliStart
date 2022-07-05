package main

type Farm struct {
	Animals []Animal
}

type Dog struct {
	Id        string `faker:"uuid_hyphenated"`
	Name      string `faker:"oneof: Sharik, Lucky, Myhtar"`
	Weight    int    `faker:"oneof: 25, 19, 30"`
	FeedPerKG int    `faker:"oneof: 5"`
}

func (d Dog) getFeedPerMonth() int {
	return d.Weight * d.FeedPerKG
}

type Cat struct {
	Id        string `faker:"uuid_hyphenated"`
	Name      string `faker:"oneof: Myrka, Chernyshka, Tigra"`
	Weight    int    `faker:"oneof: 5, 4, 3"`
	FeedPerKG int    `faker:"oneof: 7"`
}

func (c Cat) getFeedPerMonth() int {
	return c.Weight * c.FeedPerKG
}

type Cow struct {
	Id        string `faker:"uuid_hyphenated"`
	Name      string `faker:"oneof: Zorka, Raybka, Mashka"`
	Weight    int    `faker:"oneof: 300, 200, 250"`
	FeedPerKG int    `faker:"oneof: 25"`
}

func (cw Cow) getFeedPerMonth() int {
	return cw.Weight * cw.FeedPerKG
}
