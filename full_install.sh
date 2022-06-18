for cmdPath in cmd/*/; do
  echo "installing $cmdPath"
  go install ./$cmdPath
done
