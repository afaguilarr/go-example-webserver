if [ $(go fmt ./...) ]; then
  echo "Formatting was necessary"
  exit 1
else
  echo "No formatting needed" 
  exit 0
fi
