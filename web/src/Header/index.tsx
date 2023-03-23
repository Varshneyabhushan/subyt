
import { styled, alpha } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import SearchIcon from '@mui/icons-material/Search';
import HomeIcon from '@mui/icons-material/Home';
import Button from "@mui/material/Button";
import { useState } from 'react';
import { TextField, useTheme } from '@mui/material';

const Search = styled('div')(({ theme }) => ({
  position: 'relative',
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  '&:hover': {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  width: '100%',
  [theme.breakpoints.up('sm')]: {
    marginLeft: theme.spacing(10),
    width: 'auto',
  },
}));

const StyledInputBase = styled(TextField)(({ theme }) => ({
  color: 'inherit',
  '& .MuiInputBase-input': {
    padding: theme.spacing(1, 0, 1, 1),
    // vertical padding + font size from searchIcon
    transition: theme.transitions.create('width'),
    width: '100%',
    [theme.breakpoints.up('md')]: {
      width: '50ch',
    },
  },
}));

interface HeaderProps {
  initSearch : (searchTerm : string) => void;
  goHome : () => void;
}

export default function Header({ initSearch, goHome } : HeaderProps) {
  const [searchTerm, setSearchTerm] = useState("")
  const theme = useTheme()
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography
            variant="h6"
            noWrap
            component="div"
            sx={{ display: { xs: 'none', sm: 'block', fontFamily : "monaco", color : "red", fontWeight : 900 } }}
          >
            SubYT
          </Typography>
          <Typography
            variant="subtitle2"
            noWrap
            component="div"
            sx={{ margin : "10px",  display: { xs: 'none', sm: 'block' }, color : theme.palette.primary.main }}
          >
            {process.env.REACT_APP_TOPIC}
          </Typography>
          <Button variant='text' sx={{ marginLeft : 1 }} onClick={() => goHome()}>
            <HomeIcon/>
          </Button>
          <Search>
            <StyledInputBase
            variant='standard'
              value={searchTerm}
              onChange={e => setSearchTerm(e.target.value)}
              placeholder="search.."
            />
            <Button variant="text" sx={{ marginLeft : 1 }} onClick={() => initSearch(searchTerm)} >
              <SearchIcon />
            </Button>
              
          </Search>
        </Toolbar>
      </AppBar>
    </Box>
  );
}