package fetcher

type Fetcher struct {
}

func NewFetcher() Fetcher {
	return Fetcher{}
}

func (f *Fetcher) Unfetch() {

}

func (f *Fetcher) IsFetching() bool {
	return false
}

func (f *Fetcher) Fetch() {

}
