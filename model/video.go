type Video struct {
	id				 int64
	title			 string
	link			 string
	chapters		 []*Chapter
	facultyIds		 []int64
	lecturedOn time.Time
	videoLength time.Duration
	Language		 string
}
