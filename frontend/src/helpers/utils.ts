export function handleValidate(_object: Record<string, string>): Record<string, string> {
    const errors: Record<string, string> = {};
    Object.keys(_object).map(key => {
        if (!_object[key]) {
            errors[key] = "required field"
        }
    });

    return errors
}
