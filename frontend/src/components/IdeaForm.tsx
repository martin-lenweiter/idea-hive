import { useState } from 'react';
import { TextInput, Textarea, Button, Container, Title, Box } from '@mantine/core';
import { createIdea } from '../api/endpoints/ideas';

const IdeaForm = () => {
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
    <Container size="sm" mt="xl">
      <Title order={2} mb="md">
        Submit an Idea
      </Title>
      <Box component="form" onSubmit={handleSubmit}>
        <TextInput
          label="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          mb="md"
        />
        <Textarea
          label="Description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
          minRows={4}
          mb="md"
        />
        <Button
          type="submit"
          loading={loading}
        >
          {loading ? 'Submitting...' : 'Submit'}
        </Button>
      </Box>
    </Container>
  );
};

export default IdeaForm;