package paganation

var (
	DefaultPage = 1
	DefaultSize = 20
	MaxSize     = 1000
)

// 临时分页结构体
type pagination struct {
	CurrentPage uint32
	PrePage     uint32
	NextPage    uint32
	PageSize    uint32
	TotalCount  uint32
	TotalPage   uint32
	First       bool
	Last        bool
}

// NewCommonPagination 新增通用分布结构体
func NewCommonPagination(listLen, pageNo, pageSize, total uint32) *pagination {
	var first, last, hasNext bool
	var prePage, nextPage, totalCount, totalPage uint32

	if pageNo <= 1 {
		pageNo = 1
		first = true
	}

	prePage = pageNo
	if !first {
		prePage = pageNo - 1
	}

	if total > 0 {
		totalCount = total
		hasNext = totalCount > pageNo*pageSize
	} else {
		totalCount = (pageNo-1)*pageSize + listLen
		hasNext = listLen > pageSize
	}

	nextPage = pageNo
	if hasNext {
		nextPage = pageNo + 1
	}

	if pageSize != 0 {
		totalPage = totalCount / pageSize
		if totalCount%pageSize != 0 {
			totalPage += 1
		}
	}

	last = pageNo >= totalPage

	return &pagination{
		CurrentPage: pageNo,
		PrePage:     prePage,
		NextPage:    nextPage,
		PageSize:    pageSize,
		TotalCount:  totalCount,
		TotalPage:   totalPage,
		First:       first,
		Last:        last,
	}
}

// GetStartPage 获取起始查询位置
func GetStartPage(page, size, total int) int {
	if total == 0 {
		return 0
	}

	if size < 1 {
		size = 1
	}

	lastPage := total / size

	if total%size != 0 {
		lastPage++
	}

	if page > lastPage {
		page = lastPage
	}

	page = (page - 1) * size

	if page < 0 {
		page = 0
	}

	return page
}

// VerifyPage 校验分页参数
func VerifyPage(page, size int) (rtPage int, rtSize int) {
	if page <= 0 {
		rtPage = DefaultPage
	} else {
		rtPage = page
	}
	if size == 0 {
		rtSize = DefaultSize
	} else if size > MaxSize {
		rtSize = MaxSize
	} else {
		rtSize = size
	}
	return
}
