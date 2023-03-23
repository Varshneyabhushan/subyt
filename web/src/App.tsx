
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Container, CssBaseline, Pagination, Stack, useStepContext } from '@mui/material';
import Header from './Header';
import VideosList from './VideosList';
import { Suspense, useEffect, useState } from 'react';
import VideosService, { VideosServiceConfig, VideoResource } from './services/videos';
import ErrorBoundary from './Components/ErrorBoundary';
import videosResource from './services/videosResource';
import { FmdBadTwoTone } from '@mui/icons-material';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const config: VideosServiceConfig = { apiUrl: "http://localhost:3001" }
const videosService = new VideosService(config)

const videosPerPage = 20
const initialVideos = videosService.getVideos(0, videosPerPage)

function App() {

  const [resource, setResource] = useState<VideoResource>(initialVideos)
  const [videosCount, setVideosCount] = useState(100)
  const [page, setPage] = useState(1)
  const [searchTerm, setSearchTerm] = useState("")

  //false for home, true for search
  const [currentContext, setCurrentContext] = useState(false)

  useEffect(() => {
    let skip = (page - 1)* videosPerPage

    //home page vs search page
    let newResource = (currentContext) ? 
      videosService.searchVideos(searchTerm, skip, videosPerPage) : 
      videosService.getVideos(skip, videosPerPage);

    setResource(newResource)
  },
  [page, setResource, searchTerm, currentContext])

  function initSearch(term : string) {
    setSearchTerm(term)
    setPage(1)
    setCurrentContext(true)
  }

  function goHome() {
    setCurrentContext(false)
    setPage(1)
  }

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <Header initSearch={initSearch} goHome={goHome}/>
      <Pagination 
        count={Math.ceil(videosCount/20)} 
        page={page} 
        onChange={(_, pageNumber) => setPage(pageNumber)} 
        shape="rounded"/>
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
