
import ImageList from '@mui/material/ImageList';
import ImageListItem from '@mui/material/ImageListItem';
import ImageListItemBar from '@mui/material/ImageListItemBar'

interface Video {
  thumbnail : string;
  title : string;
  author : string;
}

interface CustomImageListProps {
  videos : Video[];
}

export default function VideosList({ videos } : CustomImageListProps) {
  return (
    <ImageList cols={4} gap={5} sx={{ paddingLeft : "15px", paddingRight : "15px" }}>
      {videos.map((item) => (
        <ImageListItem key={item.thumbnail}>
          <img
            src={`${item.thumbnail}?w=248&fit=crop&auto=format`}
            srcSet={`${item.thumbnail}?w=248&fit=crop&auto=format&dpr=2 2x`}
            alt={item.title}
            loading="lazy"
          />
          <ImageListItemBar
            title={item.title}
            subtitle={<span>{item.author}</span>}
            position="below"
          />
        </ImageListItem>
      ))}
    </ImageList>
  );
}