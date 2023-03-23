
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Container, CssBaseline } from '@mui/material';
import Header from './Header';
import VideosList from './VideosList';
import { Suspense, useEffect, useState } from 'react';
import VideosService, { VideosServiceConfig, VideoResource, CountResource } from './services/videos';
import ErrorBoundary from './Components/ErrorBoundary';
import PaginationCounter from './PaginationCounter';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

const config: VideosServiceConfig = { apiUrl: process.env.REACT_APP_SERVICE_URL ?? "" }
const videosPerPage = 20

console.log('apiUrl is : ', config.apiUrl)

const videosService = new VideosService(config)


const initialVideos = videosService.getVideos(0, videosPerPage)
const videosCount = videosService.getVideosCount()

function App() {

  const [videoResource, setVideoResource] = useState<VideoResource>(initialVideos)
  const [countResource, setCountResource] = useState<CountResource>(videosCount)

  const [page, setPage] = useState(1)
  const [searchTerm, setSearchTerm] = useState("")

  //false for home, true for search
  const [currentContext, setCurrentContext] = useState(false)

  useEffect(() => {
    let skip = (page - 1) * videosPerPage
    //home page
    if (currentContext) {
      let [countResource, videosResource] = videosService.searchVideos(searchTerm, skip, videosPerPage)
      setCountResource(countResource)
      setVideoResource(videosResource)
      return
    }

    //search page
    let videoResource = videosService.getVideos(skip, videosPerPage);
    setVideoResource(videoResource)
  },
    [page, setVideoResource, searchTerm, currentContext])

  function initSearch(term: string) {
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
      <Header initSearch={initSearch} goHome={goHome} />
      <ErrorBoundary fallback={"error loading count"}>
        <Suspense fallback={"loading..."}>
          <PaginationCounter page={page} setPage={setPage} countResource={countResource} />
        </Suspense>
      </ErrorBoundary>
      <Container disableGutters sx={{
        overflow: "auto",
        maxHeight: "84vh",
      }}>
        <ErrorBoundary fallback={"error while loading videos"}>
          <Suspense fallback={<h1> Loading </h1>}>
            <VideosList resource={videoResource} />
          </Suspense>
        </ErrorBoundary>
      </Container>
    </ThemeProvider>
  );
}

export default App;
