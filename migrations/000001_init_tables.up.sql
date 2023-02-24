CREATE SEQUENCE IF NOT EXISTS users_seq;

CREATE TABLE IF NOT EXISTS users
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    internal_id     INT NOT NULL,
    name            TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    module_progress INT NOT NULL DEFAULT 1,
    english_level   TEXT NOT NULL DEFAULT 'A0',
    word_count      INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS eng_words_1
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    word            TEXT NOT NULL,
    transcription   TEXT NOT NULL ,
    module          INT NOT NULL
);


CREATE TABLE IF NOT EXISTS rus_words_1
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    word            TEXT NOT NULL,
    eng_id          uuid NOT NULL ,
    module          INT NOT NULL ,
    FOREIGN KEY (eng_id) REFERENCES eng_words_1(id)
);


CREATE TABLE IF NOT EXISTS sentences_1
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    eng             TEXT NOT NULL,
    rus             TEXT NOT NULL,
    module          INT NOT NULL
);

CREATE TABLE IF NOT EXISTS progress
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL,
    module          INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)

);


CREATE TABLE IF NOT EXISTS progress_module_1
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL,
    sentences_id    UUID NOT NULL,
    answers         INT NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (sentences_id) REFERENCES sentences_1(id)
);


INSERT INTO eng_words_1(word, transcription, module) VALUES
                                                         ( 'English', '`ingliʃ', 1),
                                                         ( 'Speak', 'spi:k', 1),
                                                         ( 'There', 'ðeə', 1),
                                                         ( 'Here', 'hiə', 1),
                                                         ( 'America', 'ə`merikə', 1),
                                                         ( 'Book', 'bʊk', 1),
                                                         ( 'Know', 'nəʊ', 1),
                                                         ( 'I', 'aı', 1),
                                                         ( 'See', 'si:', 1),
                                                         ( 'Understand', 'ˌʌndəˈstænd', 1),
                                                         ( 'Agree', 'ə`gri:', 1),
                                                         ( 'Study', '`stʌdi', 1),
                                                         ( 'Work', 'wɜ:k', 1),
                                                         ( 'Remember', 'rı`membə', 1),
                                                         ( 'Read', 'ri:d', 1),
                                                         ( 'Write', 'raıt', 1),
                                                         ( 'We', 'wi:', 1),
                                                         ( 'You', 'ju:', 1),
                                                         ( 'They','ðeı',  1),
                                                         ('Message', 'ˈmesɪdʒ', 1);

INSERT INTO rus_words_1("word", "eng_id", "module" ) VALUES
                                                         ('Английский', (SELECT id FROM eng_words_1 WHERE word = 'English'), 1),
                                                         ('Говорить', (SELECT id FROM eng_words_1 WHERE word = 'Speak'), 1),
                                                         ('Там', (SELECT id FROM eng_words_1 WHERE word = 'There'), 1),
                                                         ('Здесь', (SELECT id FROM eng_words_1 WHERE word = 'Here'), 1),
                                                         ('Америка', (SELECT id FROM eng_words_1 WHERE word = 'America'), 1),
                                                         ('Книга', (SELECT id FROM eng_words_1 WHERE word = 'Book'), 1),
                                                         ('Знать', (SELECT id FROM eng_words_1 WHERE word = 'Know'), 1),
                                                         ('Я', (SELECT id FROM eng_words_1 WHERE word = 'I'), 1),
                                                         ('Видеть', (SELECT id FROM eng_words_1 WHERE word = 'See'), 1),
                                                         ('Понимать', (SELECT id FROM eng_words_1 WHERE word = 'Understand'), 1),
                                                         ('Соглашаться', (SELECT id FROM eng_words_1 WHERE word = 'Agree'), 1),
                                                         ('Учиться', (SELECT id FROM eng_words_1 WHERE word = 'Study'), 1),
                                                         ('Работать', (SELECT id FROM eng_words_1 WHERE word = 'Work'), 1),
                                                         ('Помнить', (SELECT id FROM eng_words_1 WHERE word = 'Remember'), 1),
                                                         ('Читать', (SELECT id FROM eng_words_1 WHERE word = 'Read'), 1),
                                                         ('Писать', (SELECT id FROM eng_words_1 WHERE word = 'Write'), 1),
                                                         ('Мы', (SELECT id FROM eng_words_1 WHERE word = 'We'), 1),
                                                         ('Ты', (SELECT id FROM eng_words_1 WHERE word = 'You'), 1),
                                                         ('Они', (SELECT id FROM eng_words_1 WHERE word = 'They'), 1),
                                                         ('Сообщение', (SELECT id FROM eng_words_1 WHERE word = 'Message'), 1);

INSERT INTO sentences_1(eng, rus,  module) VALUES
                                               ('The live in America', 'Они живут в Америке', 1),
                                               ('I agree with you', 'Я согласен с тобой', 1),
                                               ('I remember', 'Я помню', 1),
                                               ('I write a message', 'Я пишу сообщение', 1),
                                               ('I live here', 'Я живу здесь', 1),
                                               ('We agree', 'Мы согласны', 1),
                                               ('I understand you very well', 'Я понимаю тебя очень хорошо', 1),
                                               ('I live in this country', 'Я живу в этой стране', 1),
                                               ('You understand', 'Ты понимаешь', 1),
                                               ('I read', 'Я читаю', 1),
                                               ('I speak English and Russian', 'Я говорю по английски и по русски', 1),
                                               ('You know', 'Ты знаешь', 1),
                                               ('We remember', 'Мы помним', 1),
                                               ('They understand', 'Они понимают', 1),
                                               ('We watch different movies', 'мы смотрим различные фильмы', 1),
                                               ('I understand', 'Я понимаю', 1),
                                               ('I agree with it.', 'Я согласин с этим', 1),
                                               ('I read different books', 'Я ччитаю различные книги', 1),
                                               ('I study here', 'Я цчцсь здесь', 1),
                                               ('We work', 'Мы работаем', 1),
                                               ('I speak Russian and English', 'Я говорю по русски и по английски', 1),
                                               ('I live in Russia', 'Я живу в России', 1),
                                               ('We read', 'мы читаем', 1),
                                               ('They write a message', 'Они пишут сообщение', 1),
                                               ('I write a message', 'Я пишу сообщение', 1),
                                               ('They study', 'они учатся', 1),
                                               ('I live in this city', 'Я живу в этом городе', 1),
                                               ('I see it now', 'Я вижу это сейчас', 1),
                                               ('They live in this country', 'Они живут в этой стране', 1),
                                               ('We live in the city', 'Мы живем в этом городе', 1),
                                               ('I see it', 'Я вижу это', 1),
                                               ('I agree', 'Я согласен', 1),
                                               ('I study there', 'Я учусь там', 1),
                                               ('They agree', 'Они согласны', 1),
                                               ('I know it', 'Я знаю это', 1),
                                               ('I study', 'Я учусь', 1),
                                               ('I work', 'Я работаю', 1),
                                               ('I see', 'Я вижу', 1),
                                               ('I know it', 'Я знаю это', 1),
                                               ('I know it very well', 'Я знаю это очень хорошо', 1),
                                               ('I remember', 'Я помню', 1),
                                               ('They know it', 'Они знают это', 1),
                                               ('I read', 'Я читаю', 1),
                                               ('I know', 'Я знаю', 1),
                                               ('I see you', 'Я вижу тебя', 1),
                                               ('I speak Russian', 'Я говорю по русски', 1),
                                               ('We understand', 'Мы понимаем', 1),
                                               ('I remember it', 'Я помню это', 1),
                                               ('I understand it', 'Я понимаю это', 1);