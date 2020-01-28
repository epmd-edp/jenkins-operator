package consts

const (
	ImportStrategy = "import"

	//statuses
	StatusFailed     = "failed"
	StatusFinished   = "created"
	StatusInProgress = "in progress"

	JenkinsKind  = "Jenkins"
	CodebaseKind = "Codebase"
	StageKind    = "Stage"

	AuthorizationApiGroup = "rbac.authorization.k8s.io"
	ClusterRoleKind       = "ClusterRole"

	JenkinsServiceAccount         = "jenkins"
	EdpAdminConsoleServiceAccount = "edp-admin-console"

	LibraryCodebase = "library"
)
