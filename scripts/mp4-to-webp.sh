
INFILE=$1
OUTFILE=$(echo $INFILE | sed 's/\.[^.]*$//').webp
ffmpeg -i $INFILE -vcodec libwebp -filter:v fps=fps=15 -lossless 0 -compression_level 6 -q:v 50 -loop 0 -preset picture -an -vsync 0 $OUTFILE

