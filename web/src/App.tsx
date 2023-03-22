
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Container, CssBaseline } from '@mui/material';
import Header from './Header';
import VideosList from './VideosList';
import { Suspense, useEffect, useState } from 'react';
import VideosService, { VideosServiceConfig, Video, VideoResource } from './services/videos';
import ErrorBoundary from './Components/ErrorBoundary';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const config : VideosServiceConfig = { apiUrl : "http://localhost:3000" } 
const videosService = new VideosService(config)

function App() {

  const [resource, setResource] = useState<VideoResource>(videosService.getVideos(0, 20))
  
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <Header />
      <Container disableGutters sx={{
        overflow: "auto",
        maxHeight: "91vh",
      }}>
        <Suspense fallback={ <h1> Loading </h1>}>
        <ErrorBoundary fallback={"error while loading videos"}>
            <VideosList resource={resource} />
        </ErrorBoundary>
        </Suspense>
      </Container>
    </ThemeProvider>
  );
}

export default App;
