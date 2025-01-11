package model

const (
	BioAuthKindFace   = 1 // 人脸
	BioAuthKindVoice  = 2 // 声纹
	BioAuthKindFinger = 3 // 指纹
	BioAuthKindVein   = 4 // 静脉
	BioAuthKindPalm   = 5 // 掌纹
	BioAuthKindIris   = 6 // 虹膜
)

// UserBioAuth 生物验证 (UIndex = uicid) (TODO:GG 只存是否验证过，就没必要接口了)
type UserBioAuth struct {
	*Base
	UICId int64 `json:"uicid"` // 用户身份证id (一对一)

	Auths  map[int]bool                   `json:"auths"`  // [kind]是否认证过
	AuthAt map[int]string                 `json:"authAt"` // [kind]验证时间
	Extra  map[int]map[string]interface{} `json:"extra"`  // [kind]扩展 (以防存生物特征)
}

func NewUserBioAuth(
	base *Base,
	uicid int64,
	auths map[int]bool,
	authAt map[int]string,
	extra map[int]map[string]interface{},
) *UserBioAuth {
	return &UserBioAuth{
		Base:   base,
		UICId:  uicid,
		Auths:  auths,
		AuthAt: authAt,
		Extra:  extra,
	}
}

func (b *UserBioAuth) IsAuthed() ([]int, bool) {
	var kinds []int
	for v, ok := range b.Auths {
		if ok {
			kinds = append(kinds, v)
		}
	}
	return kinds, len(kinds) > 0
}

func (b *UserBioAuth) GetAuthByKind(kind int) bool {
	if v, ok := b.Auths[kind]; ok {
		return v
	}
	return false
}

func (b *UserBioAuth) GetAuthAtByKind(kind int) string {
	if v, ok := b.AuthAt[kind]; ok {
		return v
	}
	return ""
}

func (b *UserBioAuth) GetExtraByKind(kind int) map[string]interface{} {
	if v, ok := b.Extra[kind]; ok {
		return v
	}
	return map[string]interface{}{}
}
