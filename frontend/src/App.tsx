import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { MantineProvider, createTheme } from '@mantine/core';
import { useColorScheme } from '@mantine/hooks';
import Header from './components/Header';
import IdeaForm from './components/IdeaForm';
import Manifesto from './pages/Manifesto';

const theme = createTheme({
  primaryColor: 'blue',
  defaultRadius: 'md',
});

const App: React.FC = () => {
  const colorScheme = useColorScheme();

  return (
    <MantineProvider defaultColorScheme={colorScheme} theme={theme}>
      <Router>
        <Header />
        <Routes>
          <Route path="/" element={<IdeaForm />} />
          <Route path="/manifesto" element={<Manifesto />} />
        </Routes>
      </Router>
    </MantineProvider>
  );
};

export default App;