dirs=$(find . -type d -mindepth 1 -maxdepth 1)  # List directories in root
matrix_entries=()
for dir in $dirs; do
    if [[ -f "$dir/go.mod" ]]; then
        matrix_entries+=("\"${dir#./}\"")
    fi
done
if [ ${#matrix_entries[@]} -gt 0 ]; then
    echo "::set-output name=package-dirs::[${matrix_entries[*]}]"
else
    echo "::set-output name=package-dirs::[]"
fi