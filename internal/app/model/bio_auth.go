package model

const (
	BioAuthKindFace   = 1 // 人脸
	BioAuthKindVoice  = 2 // 声纹
	BioAuthKindFinger = 3 // 指纹
	BioAuthKindVein   = 4 // 静脉
	BioAuthKindPalm   = 5 // 掌纹
	BioAuthKindIris   = 6 // 虹膜
)

// BioAuth 生物验证 (UIndex = bid) (TODO:GG 只存是否验证过，就没必要接口了)
type BioAuth struct {
	*Base
	Bid    int64                     `json:"bid"`    // 身份卡id (一对一)
	Auths  map[int]bool              `json:"auths"`  // 认证集合
	AuthAt map[int]string            `json:"authAt"` // 验证时间
	Extra  map[int]map[string]string `json:"extra"`  // 扩展 (以防存生物特征)
}

func NewBioAuth(
	base *Base,
	bid int64,
	auths map[int]bool,
	authAt map[int]string,
	extra map[int]map[string]string,
) *BioAuth {
	return &BioAuth{
		Base:   base,
		Bid:    bid,
		Auths:  auths,
		AuthAt: authAt,
		Extra:  extra,
	}
}

func (b *BioAuth) IsAuth() (int, bool) {
	for v, ok := range b.Auths {
		if ok {
			return v, ok
		}
	}
	return -1, false
}

func (b *BioAuth) AuthByKind(kind int) bool {
	if v, ok := b.Auths[kind]; ok {
		return v
	}
	return false
}

func (b *BioAuth) AuthAtByKind(kind int) string {
	if v, ok := b.AuthAt[kind]; ok {
		return v
	}
	return ""
}

func (b *BioAuth) ExtraByKind(kind int) map[string]string {
	if v, ok := b.Extra[kind]; ok {
		return v
	}
	return map[string]string{}
}
