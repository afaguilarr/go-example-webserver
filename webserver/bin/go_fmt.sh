if $(go fmt ./main); then
  go fmt ./main
  echo "Format was necessary"
  exit 1
else
  echo "No format needed" 
  exit 0
fi
