import React, { useState } from 'react';
import { TextField, Button, Container, Typography, Box } from '@mui/material';
import { createIdea } from '../api/endpoints/ideas';

const IdeaForm: React.FC = () => {
  // Local state for form fields
  const [title, setTitle] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  // Handle form submission
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault(); // Prevent default form submission behavior
    setLoading(true);

    // Data to send in the POST request
    const idea = {
      title,
      description,
    };

    try {
      // POST request to backend
      await createIdea(idea);
      alert('Idea submitted successfully!');
      // Reset form fields after successful submission
      setTitle('');
      setDescription('');
    } catch (error) {
      console.error('Failed to submit idea:', error);
      alert('Error submitting the idea.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Container sx={{ paddingTop: 4 }}>
      <Typography variant="h4" gutterBottom>
        Submit an Idea
      </Typography>
      <Box component="form" onSubmit={handleSubmit}>
        <TextField
          label="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          fullWidth
          required
          margin="normal"
        />
        <TextField
          label="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          fullWidth
          required
          multiline
          rows={4}
          margin="normal"
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          disabled={loading}
        >
          {loading ? 'Submitting...' : 'Submit'}
        </Button>
      </Box>
    </Container>
  );
};

export default IdeaForm;