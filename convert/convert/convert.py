"""Welcome to Reflex! This file outlines the steps to create a basic app."""
from rxconfig import config

import reflex as rx

docs_url = "https://reflex.dev/docs/getting-started/introduction"
filename = f"{config.app_name}/{config.app_name}.py"


class State(rx.State):
    """The app state."""

    async def handle_image_upload(self, files: list[rx.UploadFile]):
        for file in files:
            upload_data = await file.read()
            outfile = f"/tmp/{file.filename}"
            print("image", outfile)

            # Save the file.
            with open(outfile, "wb") as file_object:
                file_object.write(upload_data)

    async def handle_video_upload(self, files: list[rx.UploadFile]):
        for file in files:
            upload_data = await file.read()
            outfile = f"/tmp/{file.filename}"
            print("video", outfile)

            # Save the file.
            with open(outfile, "wb") as file_object:
                file_object.write(upload_data)

    async def handle_upload(
            self, files: list[rx.UploadFile]
    ):
        """Handle the upload of file(s).

        Args:
            files: The uploaded files.
        """
        print("HERE BITCHCH")

        print(files)
        for file in files:
            upload_data = await file.read()
            outfile = f"/tmp/{file.filename}"
            print(outfile)

            # Save the file.
            with open(outfile, "wb") as file_object:
                file_object.write(upload_data)


def generate_upload_interface(title, accept, convert_title, convert_handler):
    return rx.container(
        rx.box(
            rx.vstack(
                rx.heading(title, font_size="2em"),
                rx.box(
                    rx.upload(
                        rx.text("Upload WebP File"),
                        border="1px solid rgb(107,99,246)",
                        padding="5em",
                        accept=accept
                    ),
                    width="100%",
                ),

                rx.button(convert_title, on_click=lambda: convert_handler(rx.upload_files())),
                spacing="1.5em",
            ),
            width="100%",
        ))


def interface_image() -> rx.Component:
    return generate_upload_interface(
        "webp to jpg",
        {
            "image/webp": [".webp"]
        },
        "Convert Image",
        State.handle_image_upload
    )


def interface_video() -> rx.Component:
    return generate_upload_interface(
        "media to mp4",
        {
            "video/webm": [".webm"],
            "video/mp4": [".mp4"],
            "image/gif": [".gif"],
        },
        "Convert Video",
        State.handle_video_upload
    )


def index() -> rx.Component:
    return rx.fragment(
        rx.box(
            rx.tabs(
                rx.tab_list(
                    rx.tab("image"),
                    rx.tab("video"),
                ),
                rx.tab_panels(
                    rx.tab_panel(
                        rx.flex(
                            interface_image(),
                            grow='1'
                        ),
                    ),
                    rx.tab_panel(
                        rx.flex(

                            interface_video(),
                            grow='1'
                        ),
                    ),
                ),
                is_fitted=True
            ),
            width="100%",
        ),
    )


# Add state and page to the app.
app = rx.App()
app.add_page(index)
app.compile()
