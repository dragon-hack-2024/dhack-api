-- Insert a user
INSERT INTO users (name, email, weight, birth_date) 
VALUES (name, 'johnny@example.com', 75, '2001-09-11');

-- Insert challenges
INSERT INTO challenges (name, steps, file_name, duration) 
VALUES 
  ('Beginner Challenge', '{"steps":[{"bpm":80,"time":10},{"bpm":100,"time":20}]}', 'beginner.mp3', 65),
  ('Marathon', '{"steps":[{"bpm":80,"time":15},{"bpm":100,"time":25}]}', 'intermediate.mp3', 320),
  ('Beat Challenge', '{"steps":[{"bpm":120,"time":20},{"bpm":140,"time":30}]}', 'advanced.mp3', 45),
  ('Expert Challenge', '{"steps":[{"bpm":140,"time":25},{"bpm":160,"time":35}]}', 'expert.mp3', 125);

-- Insert stats for past week with hard-coded values and day subtraction
INSERT INTO stats (calories_burned, rpm, duration, score, created_at, challenge_id, user_id)
VALUES 
  (60, 70, 60, 80, now() - INTERVAL '6 day', 1, 1),
  (120, 110, 80, 82, now() - INTERVAL '5 day', 2, 1),
  (220, 140, 120, 84, now() - INTERVAL '4 day', 3, 1),
  (20, 85, 50, 86, now() - INTERVAL '3 day', 1, 1),
  (110, 75, 150, 88, now() - INTERVAL '2 day', 2, 1),
  (80, 150, 170, 90, now() - INTERVAL '1 day', 3, 1),
  (90, 100, 55, 92, now(), 4, 1);
