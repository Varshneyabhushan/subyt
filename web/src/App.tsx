
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Container, CssBaseline } from '@mui/material';
import Header from './Header';
import VideosList from './VideosList';
import { Suspense, useState } from 'react';
import VideosService, { VideosServiceConfig, VideoResource } from './services/videos';
import ErrorBoundary from './Components/ErrorBoundary';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const config: VideosServiceConfig = { apiUrl: "http://localhost:3001" }
const videosService = new VideosService(config)

const initialVideos = videosService.getVideos(0, 20)

function App() {

  const [resource, setResource] = useState<VideoResource>(initialVideos)

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <Header />
      <Container disableGutters sx={{
        overflow: "auto",
        maxHeight: "91vh",
      }}>
        <ErrorBoundary fallback={"error while loading videos"}>
          <Suspense fallback={<h1> Loading </h1>}>
              <VideosList resource={resource} />
          </Suspense>
        </ErrorBoundary>
      </Container>
    </ThemeProvider>
  );
}

export default App;
