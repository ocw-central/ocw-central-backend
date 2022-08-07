type Subject struct {
	id				 int64 
	category		 string 
	title			 string 
	facultyIds		 []int64
	videoIds		 []int64
	location		 string
	remark 			 string
	pdfLinks		 []string
	relatedSubjectIds []int64
	department       string
	language         string
	academicYear     string
	semester         string
	firstHeldOn      string
	numCredit       int8
	courceFormat     string
	targetedAudience string
	dayOfWeek        string
	courcePeriod     string
	outline		  string
	objective	  string
	lessonPlan	  string
	gradingMethod	  string
	courceRequirement string
	reference	  string
	subpageIds	  []int64
}