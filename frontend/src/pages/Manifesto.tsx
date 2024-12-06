import React from 'react';
import { Container, Typography } from '@mui/material';
import ReactMarkdown from 'react-markdown';

const manifestoContent = `
# IdeaHive Manifesto

## Main Premise

A good idea is essential for a great project, but execution is the real challenge, as anyone who has built something great will tell you.

Therefore, by sharing your ideas, you don't lose anything, but you might gain much. You open yourself to the collective intelligence and new opportunities you might not have expected might emerge.

## Why might you want to join the idea hive?

- You need something built but can’t do it yourself.
- You want to build but need ideas.
- You have more ideas than you can execute.
- You want to contribute to the collective idea pool for the advancement of society.
- You’re seeking cofounders.
- You want to refine your idea through collective brainstorming and feedback from domain experts.
- You are looking for someone to invest in your project.
- You are looking to invest in projects.
- You want to test the market potential of your idea before committing resources.

`;

const Manifesto: React.FC = () => {
  return (
    <Container>
      <ReactMarkdown>{manifestoContent}</ReactMarkdown>
    </Container>
  );
};

export default Manifesto;