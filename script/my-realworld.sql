-- CreateTable
CREATE TABLE `User` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` TEXT NOT NULL,
    `username` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `image` Varchar(200) DEFAULT 'https://api.realworld.io/images/smiley-cyrus.jpeg',
    `bio` TEXT,
    `demo` BOOLEAN NOT NULL DEFAULT false
);

-- CreateTable
CREATE TABLE `Article` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `slug` TEXT NOT NULL,
    `title` TEXT NOT NULL,
    `description` TEXT NOT NULL,
    `body` TEXT NOT NULL,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `authorId` INT NOT NULL,
    CONSTRAINT `Article_authorId_fkey` FOREIGN KEY (`authorId`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE `Comment` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `body` TEXT NOT NULL,
    `articleId` INT NOT NULL,
    `authorId` INT NOT NULL,
    CONSTRAINT `Comment_articleId_fkey` FOREIGN KEY (`articleId`) REFERENCES `Article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `Comment_authorId_fkey` FOREIGN KEY (`authorId`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE `Tag` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL
);

-- CreateTable
CREATE TABLE `_ArticleToTag` (
    `A` INT NOT NULL,
    `B` INT NOT NULL,
    CONSTRAINT `_ArticleToTag_A_fkey` FOREIGN KEY (`A`) REFERENCES `Article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `_ArticleToTag_B_fkey` FOREIGN KEY (`B`) REFERENCES `Tag` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE `_UserFavorites` (
    `A` INT NOT NULL,
    `B` INT NOT NULL,
    CONSTRAINT `_UserFavorites_A_fkey` FOREIGN KEY (`A`) REFERENCES `Article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `_UserFavorites_B_fkey` FOREIGN KEY (`B`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE `_UserFollows` (
    `A` INT NOT NULL,
    `B` INT NOT NULL,
    CONSTRAINT `_UserFollows_A_fkey` FOREIGN KEY (`A`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `_UserFollows_B_fkey` FOREIGN KEY (`B`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateIndex
CREATE UNIQUE INDEX `Article_slug_key` ON `Article`(`slug`);

-- CreateIndex
CREATE UNIQUE INDEX `Tag_name_key` ON `Tag`(`name`);

-- CreateIndex
CREATE UNIQUE INDEX `User_email_key` ON `User`(`email`);

-- CreateIndex
CREATE UNIQUE INDEX `User_username_key` ON `User`(`username`);

-- CreateIndex
CREATE UNIQUE INDEX `_ArticleToTag_AB_unique` ON `_ArticleToTag`(`A`, `B`);

-- CreateIndex
CREATE INDEX `_ArticleToTag_B_index` ON `_ArticleToTag`(`B`);

-- CreateIndex
CREATE UNIQUE INDEX `_UserFavorites_AB_unique` ON `_UserFavorites`(`A`, `B`);

-- CreateIndex
CREATE INDEX `_UserFavorites_B_index` ON `_UserFavorites`(`B`);

-- CreateIndex
CREATE UNIQUE INDEX `_UserFollows_AB_unique` ON `_UserFollows`(`A`, `B`);

-- CreateIndex
CREATE INDEX `_UserFollows_B_index` ON `_UserFollows`(`B`);


INSERT INTO `my-realworld`.`user`(`id`, `email`, `username`, `password`, `image`, `bio`, `demo`) VALUES (1, 'haipeng_lin@163.com', '林海鹏', '123', 'https://api.realworld.io/images/smiley-cyrus.jpeg', NULL, 0);
