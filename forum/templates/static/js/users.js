class UserBlock {
    static Template = document.querySelector("#tmpl-user-block");

    static Nickname = UserBlock.Template.content.querySelector("#tmpl-user-block__nickname");
    static BlockAvatarImg = UserBlock.Template.content.querySelector("#tmpl-user-block__block-img");
    static AvatarImg = UserBlock.Template.content.querySelector("#tmpl-user-block__img");
    static CreatedTime = UserBlock.Template.content.querySelector("#tmpl-user-block__created-time");
    static QuestionsCount = UserBlock.Template.content.querySelector("#tmpl-user-block__questions-count");
    static KarmaCount = UserBlock.Template.content.querySelector("#tmpl-user-block__karma-count");

    UserName
    UserAvatarPath
    UserURL
    UserCreatedText
    UserQuestionsCount
    UserKarmaCount

    // GetCopyUserBlock - returns UserBlock as HTMLelement
    static GetCopyElementUserBlock(user) {
        UserBlock.Nickname.textContent = user.UserName
        UserBlock.Nickname.href = user.UserURL
        UserBlock.BlockAvatarImg.href = user.UserURL
        UserBlock.AvatarImg.src = user.UserAvatarPath
        UserBlock.CreatedTime.title = user.UserCreatedText
        UserBlock.CreatedTime.textContent = user.UserCreatedText
        UserBlock.QuestionsCount.textContent = user.UserQuestionsCount
        UserBlock.KarmaCount.textContent = user.UserKarmaCount

        return UserBlock.Template.content.cloneNode(true)
    }
}

const B_Users = document.querySelector(".users");

//? This is temp solution
function DownloadUsers(){
    const user = new UserBlock()
    user.UserName = "Nickname"
    user.UserAvatarPath = "static/img/avatar.jpeg"
    user.UserURL = "/"
    user.UserCreatedText = (new Date()).toISOString().split('T')[0]
    user.UserQuestionsCount = 1
    user.UserKarmaCount = 0

    AppendUserToHTML(user)
}

function AppendUserToHTML(user) {
    B_Users.append(UserBlock.GetCopyElementUserBlock(user))
}

for (let index = 0; index < 10; index++) {
    DownloadUsers()
}