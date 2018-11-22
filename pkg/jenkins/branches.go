package jenkins

import "github.com/xcloudnative/xcloud/pkg/gits"

func BranchPattern(gitKind string) string {
	switch gitKind {
	case gits.KindBitBucketCloud, gits.KindBitBucketServer:
		return BranchPatternMatchEverything
	default:
		return BranchPatternMasterPRsAndFeatures
	}
}
