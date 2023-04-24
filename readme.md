# Natural Chinese
A data-processing pipeline providing clear and concise classifications of Chinese texts available online.

## Data-flow

### Input
- New text files are read by the RSS feed parser and stored as a txt file in its original form.
- A reference to this document is saved to a queue in Redis.
- Queue workers send new files to be processed.

### Chengyu detection
- Chengyus are central to Chinese literature in all of its forms, so before we start processing the text, our first step is to identify the Chengyus in the text and mark them as special expressions on their own. Every text gets matched against a database of known Chengyus for this.
- Once the Chengyus are spotted, they are stored in Firebase as their own expressions, with a reference to their original text.

### Segmentation
- Chinese text doesn't have spaces between words so we use the Stanford NLP segmenter to read the text and separate it into words. Before running the segmenter, any Chengyus found are highlighted so that segmenting considers them a single expression.
- A de-duplicated list of all the words used in the text is then stored in Firestore, with references to how many times each word was used.
- Another version of this list, this time without stop-words is also stored.

### Classification
- The list of words of each text is then analyzed against some well known corpora of Chinese texts to infer its statistical distance from each predefined level.
- The corpora used ranges from texts generally classified as "easy" (children's books, graded readers), to "intermediate" (newspaper articles, youth books), to "advanced" (scientific articles, classic books)
- The classification of each text follows a matching pattern of general proximity.

### Output
- The statistical distances found are then translated into tags that are applied to the text, allowing each user to make a decision on whether the level of difficulty proposed is within their limits. 

## Infrastructure

### Data-storage
- Document Database (noSQL)
- Classification database (SQL)
- Worker queues database
- General storage

### Data-processing
- Functions

### Client API
- Containers
