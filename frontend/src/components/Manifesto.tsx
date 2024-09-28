import React from 'react';
import { Container, Typography } from '@mui/material';
import ReactMarkdown from 'react-markdown';

const manifestoContent = `
# Our Manifesto

Here at IdeaHive, we believe in the power of ideas to change the world...

## Our Core Values

1. **Innovation**: We embrace new ideas and technologies.
2. **Collaboration**: We believe in the power of working together.
3. **Impact**: We strive to make a positive difference in the world.

Learn more about our [mission and vision](/about).
`;

const Manifesto: React.FC = () => {
  return (
    <Container>
      <ReactMarkdown>{manifestoContent}</ReactMarkdown>
    </Container>
  );
};

export default Manifesto;