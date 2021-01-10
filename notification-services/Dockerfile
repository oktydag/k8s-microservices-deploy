FROM python AS build
RUN mkdir -p /app/notification-services
WORKDIR /app/notification-services
COPY . .
RUN pip install -r requirements.txt
CMD [ "python", "./main.py" ]
EXPOSE 5006