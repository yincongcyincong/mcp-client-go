1. `browser_close` – Closes the current browser page.

    - Close the current browser page.

2. `browser_wait` – Waits for a specified time in seconds.

    - Wait for **X** seconds.
      *(Replace **X** with the desired number of seconds)*

3. `browser_resize` – Resizes the browser window.

    - Resize the browser window to **width** x **height**.
      *(Replace **width** and **height** with the desired dimensions)*

4. `browser_console_messages` – Returns all console messages from the browser.

    - Show me the console messages from the browser.

5. `browser_handle_dialog` – Handles a dialog by accepting or rejecting it.

    - Accept or reject the dialog. If it's a prompt, provide the text for the prompt.

6. `browser_file_upload` – Uploads one or multiple files.

    - Upload the following files: **file\_path1, file\_path2, ...**.
      *(Replace **file\_path1, file\_path2, ...** with the actual file paths)*

7. `browser_install` – Installs the browser specified in the configuration.

    - Install the browser specified in the configuration.

8. `browser_press_key` – Presses a key on the keyboard.

    - Press the **key\_name** key.
      *(Replace **key\_name** with the key to be pressed, e.g., "Enter" or "ArrowLeft")*

9. `browser_navigate` – Navigates to a URL.

    - Navigate to the following URL: **url**.
      *(Replace **url** with the desired URL)*

10. `browser_navigate_back` – Goes back to the previous page.

    - Go back to the previous page.

11. `browser_navigate_forward` – Goes forward to the next page.

    - Go forward to the next page.

12. `browser_network_requests` – Lists all network requests since loading the page.

    - Show me all the network requests since loading the page.

13. `browser_pdf_save` – Saves the current page as a PDF.

    - Save the page as a PDF.

14. `browser_snapshot` – Takes an accessibility snapshot of the current page.

    - Capture an accessibility snapshot of the current page.

15. `browser_click` – Performs a click action on a web page.

    - Click on the **element\_name** on the page.
      *(Replace **element\_name** with a description of the element you wish to click on)*

16. `browser_drag` – Performs a drag and drop between two elements.

    - Drag the mouse from **start\_element** to **end\_element**.
      *(Replace **start\_element** and **end\_element** with descriptions of the source and target elements)*

17. `browser_hover` – Hovers the mouse over an element on the page.

    - Hover the mouse over **element\_name** on the page.
      *(Replace **element\_name** with the description of the element)*

18. `browser_type` – Types text into an editable element.

    - Type **text\_to\_type** into the **element\_name**.
      *(Replace **text\_to\_type** with the text to be typed, and **element\_name** with the element you want to type into)*

19. `browser_select_option` – Selects an option in a dropdown.

    - Select **value1, value2, ...** in the dropdown **element\_name**.
      *(Replace **value1, value2, ...** with the values you want to select and **element\_name** with the description of the dropdown element)*

20. `browser_take_screenshot` – Takes a screenshot of the current page.

    - Take a screenshot of **element\_name**.
      *(Replace **element\_name** with the description of the element to capture, or leave blank for the whole page)*

21. `browser_tab_list` – Lists all browser tabs.

    - Show me all the browser tabs.

22. `browser_tab_new` – Opens a new browser tab.

    - Open a new tab with the URL **url**.
      *(Replace **url** with the URL to open, or leave blank for a new blank tab)*

23. `browser_tab_select` – Selects a tab by index.

    - Select tab number **index**.
      *(Replace **index** with the index of the tab you wish to select)*

24. `browser_tab_close` – Closes a browser tab.

    - Close the tab at index **index**.
      *(Replace **index** with the index of the tab to close, or leave blank to close the current tab)*

25. `browser_generate_playwright_test` – Generates a Playwright test for a given scenario.

    - Generate a Playwright test with the name **test\_name** and description **test\_description**. The test steps are: **step1, step2, ...**.
      *(Replace **test\_name**, **test\_description**, and **step1, step2, ...** with the actual test name, description, and steps)*

