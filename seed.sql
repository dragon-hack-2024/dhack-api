-- Insert a user
INSERT INTO users (name, email, weight, height, birth_date) 
VALUES  ('Johnny', 'johnny@example.com', 75, 180, '2001-09-11'),
        ('Tony', 'tony@example.com', 85, 160, '1988-02-22'),
        ('Joe', 'joe@example.com', 70, 185, '1922-04-23');

-- Insert challenges
INSERT INTO challenges (name, steps, file_name, duration) 
VALUES 
  ('Beat Bounce', '[{"bpm":115,"time":10},{"bpm":150,"time":10},{"bpm":96,"time":10}]', 'advanced.mp3', 30),
  ('Stairway to Exhaustion', '[{"bpm":130,"time":10},{"bpm":150,"time":20}]', 'beginner.mp3', 130),
  ('Marathon', '[{"bpm":80,"time":15},{"bpm":100,"time":25}]', 'intermediate.mp3', 320),
  ('Casual', '[{"bpm":120,"time":25},{"bpm":100,"time":35}]', 'expert.mp3', 125);

-- Insert stats for past week with hard-coded values and day subtraction
INSERT INTO stats (calories_burned, rpm, duration, score, created_at, challenge_id, user_id)
VALUES 
  (60, 70, 60, 80, now() - INTERVAL '6 day', 1, 1),
  (120, 110, 80, 82, now() - INTERVAL '5 day', 2, 1),
  (220, 140, 120, 84, now() - INTERVAL '4 day', 3, 1),
  (20, 85, 50, 86, now() - INTERVAL '3 day', 1, 1),
  (110, 75, 150, 88, now() - INTERVAL '2 day', 2, 1),
  (80, 80, 150, 65, now() - INTERVAL '2 day', 1, 2),
  (85, 85, 150, 34, now() - INTERVAL '3 day', 1, 3),
  (80, 150, 170, 90, now() - INTERVAL '1 day', 3, 1),
  (90, 100, 55, 92, now(), 4, 1);
