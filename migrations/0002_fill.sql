INSERT INTO tenders (name, description, service_type, status, organization_id, creator_id) VALUES
('Software Development Project', 'Development of a new software product', 'Software Development', 'PUBLISHED', 1, 1),
('Global Marketing Campaign', 'Marketing campaign across multiple countries', 'Marketing', 'PUBLISHED', 2, 2),
('Business Consultation', 'Consulting for small business growth strategies', 'Consultation', 'PUBLISHED', 3, 3),
('Solar Power Plant Construction', 'Construction of a 50MW solar power plant', 'Construction', 'CREATED', 4, 4),
('Healthcare Mobile App', 'Development of a mobile app for healthcare services', 'Mobile Application', 'PUBLISHED', 5, 5);

INSERT INTO bids (name, description, status, tender_id, author_id, decision) VALUES
('Tech Innovations Bid', 'Bid for software development project', 'PUBLISHED', 1, 1, 'Approved'),
('Global Enterprises Marketing Bid', 'Bid for global marketing campaign', 'PUBLISHED', 2, 2, 'Under Review'),
('Smith Consulting Bid', 'Bid for business consultation services', 'PUBLISHED', 3, 3, 'Approved'),
('Green Energy Construction Bid', 'Bid for solar power plant construction', 'CREATED', 4, 4, 'Pending'),
('Healthcare Solutions App Bid', 'Bid for healthcare mobile app development', 'PUBLISHED', 5, 5, 'Rejected');

INSERT INTO reviews (bid_id, tender_id, author_username, organization_id, comment, rating) VALUES
(1, 1, 'asmith', 1, 'Excellent proposal with detailed plan.', 5),
(2, 2, 'jdoe', 2, 'Good proposal, but lacks global reach.', 4),
(3, 3, 'mjones', 3, 'Comprehensive and well thought out.', 5),
(4, 4, 'bwilliams', 4, 'Proposal is promising but needs refinement.', 3),
(5, 5, 'kthomas', 5, 'Lacks innovation and user-centric design.', 2);
