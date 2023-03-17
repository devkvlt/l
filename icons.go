package main

const (
	none = iota
	red
	green
	yellow
	blue
	magenta
	cyan
)

const (
	dirIcon         = " "
	emptyDirIcon    = " "
	defaultFileIcon = " "
)

var specialFileIcons = map[string]struct {
	icon  string
	color int8
}{
	".DS_Store":           {" ", none},
	".babelrc.json":       {" ", yellow},
	".clang-format":       {" ", none},
	".clang-tidy":         {" ", none},
	".eslintignore":       {" ", magenta},
	".eslintrc.json":      {" ", magenta},
	".git_message":        {" ", red},
	".gitconfig":          {" ", red},
	".gitignore":          {" ", red},
	".gitignore_global":   {" ", red},
	".localized":          {" ", none},
	".markdownlint.json":  {" ", red},
	".markdownlintignore": {" ", red},
	".npmrc":              {" ", red},
	".prettierignore":     {" ", magenta},
	".prettierrc.json":    {" ", magenta},
	".stylelintignore":    {" ", none},
	".stylelintrc.json":   {" ", none},
	"LICENSE":             {" ", yellow},
	"esbuild.config.js":   {" ", yellow},
	"favicon.ico":         {" ", yellow},
	"jest.config.js":      {" ", red},
	"netlify.toml":        {" ", green},
	"next.config.js":      {" ", none},
	"nodemon.json":        {" ", green},
	"package-lock.json":   {" ", red},
	"package.json":        {" ", red},
	"pnpm-lock.yaml":      {" ", yellow},
	"pnpm-workspace.yaml": {" ", yellow},
	"rollup.config.js":    {" ", red},
	"vercel.json":         {" ", none},
	"webpack.config.js":   {" ", blue},
	"yarn.lock":           {" ", blue},
}

var filetypeIcons = map[string]struct {
	icon  string
	color int8
}{
	".c":      {" ", blue},
	".cfg":    {" ", none},
	".conf":   {" ", none},
	".config": {" ", none},
	".css":    {" ", blue},
	".go":     {" ", cyan},
	".h":      {" ", blue},
	".html":   {" ", red},
	".ini":    {" ", none},
	".js":     {" ", yellow},
	".json":   {" ", none},
	".jsonc":  {" ", none},
	".jsx":    {" ", cyan},
	".lua":    {" ", blue},
	".md":     {" ", red},
	".py":     {" ", blue},
	".scss":   {" ", magenta},
	".sh":     {" ", green},
	".svg":    {" ", yellow},
	".toml":   {" ", none},
	".ts":     {" ", blue},
	".tsx":    {" ", blue},
	".vim":    {" ", green},
	".yaml":   {" ", red},
	".yml":    {" ", red},

	".txt": {" ", cyan},
	".pdf": {" ", red},
	".log": {" ", green},

	".7z":  {" ", yellow},
	".apk": {" ", yellow},
	".dmg": {" ", yellow},
	".gz":  {" ", yellow},
	".pkg": {" ", yellow},
	".rar": {" ", yellow},
	".tar": {" ", yellow},
	".zip": {" ", yellow},

	".eot":   {" ", none},
	".ttf":   {" ", none},
	".woff":  {" ", none},
	".woff2": {" ", none},
	".otf":   {" ", none},

	".bmp":  {" ", green},
	".gif":  {" ", green},
	".ico":  {" ", green},
	".jpeg": {" ", green},
	".jpg":  {" ", green},
	".png":  {" ", green},
	".tiff": {" ", green},
	".webp": {" ", green},

	".aac":  {" ", blue},
	".flac": {" ", blue},
	".m4a":  {" ", blue},
	".mp3":  {" ", blue},
	".wav":  {" ", blue},
	".wma":  {" ", blue},
	".midi": {" ", none},

	".avi":  {" ", red},
	".flv":  {" ", red},
	".m4v":  {" ", red},
	".mkv":  {" ", red},
	".mov":  {" ", red},
	".mp4":  {" ", red},
	".mpeg": {" ", red},
	".webm": {" ", red},
	".wmv":  {" ", red},
}
