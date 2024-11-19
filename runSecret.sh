# Get mail password from bash input
echo "Enter mail password: "
read -s entered_password
# Run the secret santa script

FROM="no-reply@funkemunky.cc" HOST="mail.funkemunky.cc" PORT="587" MAIL_PASSWORD=$entered_password secretSanta