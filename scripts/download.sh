#!/bin/bash
set -x

# Function to download and rename file
download_and_rename() {
    url="$1"
    temp_file=$(mktemp)
    
    # Download the file
    if curl -s "$url" -o "$temp_file"; then
        # Calculate SHA1 hash
        hash=$(sha1sum "$temp_file" | cut -d' ' -f1)
        
        # Get the first 8 characters of the hash
        short_hash=${hash:0:8}
        
        # Get the file extension from the URL
        url_without_query=$(echo "$url" | cut -d'?' -f1)
        extension="${url_without_query##*.}"
        
        # Rename the file
        mv "$temp_file" "${short_hash}.${extension}"
        echo "Downloaded and renamed: ${short_hash}.${extension}"
    else
        echo "Failed to download: $url"
        rm "$temp_file"
    fi
}

# Read URLs from stdin and process each one
while read -r url; do
    download_and_rename "$url"
done

echo "All downloads completed."
