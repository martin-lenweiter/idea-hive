import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import { CssBaseline } from '@mui/material';
import theme from './theme';
import Header from './components/Header';
import IdeaForm from './components/IdeaForm';
import Manifesto from './components/Manifesto';

const App: React.FC = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <Header />
        <Routes>
          <Route path="/" element={<IdeaForm />} />
          <Route path="/manifesto" element={<Manifesto />} />
        </Routes>
      </Router>
    </ThemeProvider>
  );
};

export default App;