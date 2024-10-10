import React from 'react';
import { AppBar, Toolbar, Typography, Button, Box } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import logo from '../logo.png';

const Header: React.FC = () => {
  return (
    <AppBar position="static" sx={{ backgroundColor: '#30253f' }}>
      <Toolbar>
        <Box sx={{ display: 'flex', alignItems: 'center', flexGrow: 1 }}>
          <img src={logo} alt="IdeaHive Logo" style={{ height: '60px', marginRight: '10px' }} />
          <Typography variant="h6" component="div">
            IdeaHive
          </Typography>
        </Box>
        <Button color="inherit" component={RouterLink} to="/">
          Home
        </Button>
        <Button color="inherit" component={RouterLink} to="/manifesto">
          Manifesto
        </Button>
      </Toolbar>
    </AppBar>
  );
};

export default Header;