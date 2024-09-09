import apiClient from '../apiClient';
import { Idea } from '../../types/idea';

export const createIdea = (idea: Idea) => {
  return apiClient.post('/ideas', idea);
};