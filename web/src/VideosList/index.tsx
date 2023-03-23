
import { Link } from '@mui/material';
import ImageList from '@mui/material/ImageList';
import ImageListItem from '@mui/material/ImageListItem';
import ImageListItemBar from '@mui/material/ImageListItemBar'
import { VideoResource } from '../services/videos';


interface CustomImageListProps {
  resource : VideoResource;
}

const youtubeVideo = (id : string) => `https://www.youtube.com/watch?v=${id}`

export default function VideosList({ resource } : CustomImageListProps) {
  const videos = resource.read()
  
  return (
    <ImageList cols={4} gap={5} sx={{ paddingLeft : "15px", paddingRight : "15px" }}>
      {videos.map((item) => (
        <ImageListItem 
        key={item.Id} 
        sx={{ width : 220, cursor : "pointer" }} 
        onClick={() => window.open(youtubeVideo(item.Id), "_blank")}>
          <img
            src={`${item.Thumbnails[0].Url}`}
            // srcSet={`${item.thumbnail}?w=248&fit=crop&auto=format&dpr=2 2x`}
            alt={item.Title}
            loading="lazy"
          />
          <ImageListItemBar
            title={item.Title}
            subtitle={<span>{item.Channel.Title}</span>}
            position="below"
          />
        </ImageListItem>
      ))}
    </ImageList>
  );
}