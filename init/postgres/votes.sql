-- Create the voting database
CREATE DATABASE voting;

-- Connect to the voting database
\c voting

-- Create the votes table
-- This table stores individual votes with a unique ID, survey identifier, question number, and a timestamp of when the vote was created
CREATE TABLE votes (
  id TEXT PRIMARY KEY,    -- Unique identifier for the vote
  survey TEXT,            -- Identifier for the survey
  question INT,           -- Number of the question being voted on
  created BIGINT          -- Timestamp of when the vote was created
);

-- Create the results table
-- This table stores the total vote count for each survey and question, along with the timestamp of the last update
CREATE TABLE results (
  survey TEXT,            -- Identifier for the survey
  question INT,           -- Number of the question being voted on
  votes INT,              -- Total number of votes for the question
  last_update BIGINT,     -- Timestamp of the last update
  PRIMARY KEY(survey, question)  -- Primary key consisting of the survey and question combination
);
