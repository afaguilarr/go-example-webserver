files=$(go fmt ./main)
if files; then
  echo files
  echo "Format was necessary"
  exit 1
else
  echo "No format needed" 
  exit 0
fi
